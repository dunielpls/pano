#!/usr/bin/env python3
# -*- coding: utf-8; mode: python; tab-width: 4 -*-
# vim: ft=python fenc=utf-8
#
# Author:  Daniel Isaksen <daniel.isaksen@banenor.no>
#

from typing import Union, Optional

import asyncio
import aiohttp
import json

class APIException(Exception):
    def __init__(self, *args, **kwargs):
        super().__init__(*args)

        self.error = kwargs.get("error", None)

class Zabbix(object):
    def __init__(self,
                 url:      str,
                 username: str,
                 password: str):
        self._url      = url
        self._username = username
        self._password = password

        # Initialized during login.
        self._token = None

    async def _aio_startup(self):
        # `ClientSession` has to be instantiated inside a coroutine.
        self._session = aiohttp.ClientSession()

        self._session.headers.update({
            "Content-Type":  "application/json-rpc",
            "User-Agent":    "Pano/0.0.1",
            "Cache-Control": "no-cache"
        })

    async def _aio_shutdown(self):
        await self._session.close()

    async def _login(self):
        response = await self.user.login(
            user      = self._username,
            password  = self._password,
            _is_login = True
        )

        if "error" in response:
            raise APIException(
                response["error"]["message"],
                error = response["error"]["code"]
            )

        self._token = response.get("result")

    async def do(self,
                 method:    str,
                 params:    dict,
                 login: bool = False):
        # No pre-established session, perform login.
        if not login and not self._token:
            await self._login()

        request_params = {
            "jsonrpc": "2.0",
            "id":      1,
            "auth":    self._token,
            "method":  method,
            "params":  params
        }

        response_obj = await self._session.post(
            url  = self._url,
            json = request_params
        )

        response = await response_obj.json()

        if "error" in response:
            if response["error"].get("code", 0) == "-32602":
                # Not logged in - perform login.
                pass#await self._login()
            else:
                raise APIException(
                    response["error"]["message"],
                    error = response["error"]["code"]
                )

        # TODO: Extract actual data / error?
        return response

    def __getattr__(self, first_part: str):
        # With the example API method "host.get", this function handes the
        # "host" part, then returns a wrapper class that handles "get".
        return _MethodWrapper(first_part, self)

class _MethodWrapper(object):
    def __init__(self, first_part: str, parent: Zabbix):
        self._first_part = first_part
        self._parent = parent

    def __getattr__(self, second_part: str):
        # Return a function that handles the method parameters.
        def fn(**kwargs):
            # Default to `False`,
            login = False

            # Extract the '_is_login' parameter for the method 'user.login'.
            if self._first_part == "user" and second_part == "login":
                login = kwargs.pop("_is_login", False)

            return self._parent.do(
                method = f"{self._first_part}.{second_part}",
                params = kwargs,
                login  = login
            )

        return fn

if __name__ == "__main__":
    z = Zabbix(
        url      = "http://10.10.10.100/api_jsonrpc.php",
        username = "sa_pano",
        password = "sa_pano"
    )

    async def main_test():
        await z._init_aio()
        await z.host.get()

    asyncio.run(main_test())
