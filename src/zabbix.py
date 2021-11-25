#!/usr/bin/env python3
# -*- coding: utf-8; mode: python; tab-width: 4 -*-
# vim: ft=python fenc=utf-8
#
# Author:  Daniel Isaksen <daniel.isaksen@banenor.no>
#

import asyncio
import aiohttp

class APIException(Exception):
    def __init__(self, *args, **kwargs):
        super().__init__(*args)

        # Grab the error code. TODO: Grab more information?
        self.code = kwargs.get("code", None)

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

        # Set required headers up front.
        self._session.headers.update({
            "Content-Type":  "application/json-rpc",
            "User-Agent":    "Pano/0.0.1",
            "Cache-Control": "no-cache"
        })

    async def _aio_shutdown(self):
        # Cleanly shut down the aiohttp session or else it complains. Loudly.
        await self._session.close()

    async def _login(self):
        # Call the API method `user.login`.
        response = await self.user.login(
            user      = self._username,
            password  = self._password,
            _is_login = True
        )

        # Oops.
        if "error" in response:
            raise APIException(
                response["error"]["message"],
                code = response["error"]["code"]
            )

        # Extract the token itself.
        self._token = response.get("result")

    async def do(self,
                 method:    str,
                 params:    dict,
                 _is_login: bool = False):
        # No pre-established session, perform login. The `not login` part is to
        # avoid calling `_login` again before authentication has finished.
        if not _is_login and not self._token:
            await self._login()

        # TODO: Don't send an API token for certain calls like `apiinfo.version`
        # and `user.checkAuthentication`.

        # Some of these fields are required by the Zabbix API.
        request_params = {
            "jsonrpc": "2.0",
            "id":      "_",         # Arbitrary identifier.
            "auth":    self._token, # Authentication token
            "method":  method,
            "params":  params
        }

        # Make the underlying HTTP call.
        response_obj = await self._session.post(
            url  = self._url,
            json = request_params
        )

        # Extract JSON from the response.
        response = await response_obj.json()

        # Handle errors.
        if "error" in response:
            # The following code means we are not logged in.
            if response["error"].get("code", 0) == "-32602":
                await self._login()
            else:
                # TODO: Format the exception message in a clearer way?
                raise APIException(
                    response["error"]["message"],
                    code = response["error"]["code"]
                )

        return response

    # If you call an instance of this class to execute an API method, say
    # `zbx.host.get`, this function is called as `zbx.__getattr__("host")`.
    def __getattr__(self, first_part: str):
        # Return a wrapper class that handles the `get` part.
        return _MethodWrapper(first_part, self)

class _MethodWrapper(object):
    def __init__(self, first_part: str, parent: Zabbix):
        self._first_part = first_part
        self._parent = parent

    # Similar to `Zabbix.__call__("host")`, this function is called as
    # `_MethodWrapper.__getattr__("get")`
    def __getattr__(self, second_part: str):
        # Return a function that handles the method parameters.
        def fn(**kwargs):
            # Default to `False`,
            _is_login = False

            # Extract the '_is_login' parameter for the method 'user.login'.
            if self._first_part == "user" and second_part == "login":
                _is_login = kwargs.pop("_is_login", False)

            # Call the parent's `do` function to do the actual work.
            return self._parent.do(
                method     = f"{self._first_part}.{second_part}",
                params     = kwargs,
                _is_login  = _is_login
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
