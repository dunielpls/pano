#!/usr/bin/env python3
# -*- coding: utf-8; mode: python; tab-width: 4 -*-
# vim: ft=python fenc=utf-8
#
# Author:  Daniel Isaksen <daniel.isaksen@banenor.no>
#

from typing import Optional, List

from fastapi             import (FastAPI, Query)
from fastapi.routing     import APIRouter
from fastapi.staticfiles import StaticFiles

from zabbix import Zabbix

app = FastAPI()
api = APIRouter()

@app.on_event("startup")
async def startup_event():
    z = Zabbix(
        url      = "http://10.10.10.100/api_jsonrpc.php",
        username = "sa_pano",
        password = "sa_pano"
    )

    # Initialize asyncio stuff. TODO: Better way to do this?
    await z._aio_startup()

    # Make the API client available through `app`.
    app.zabbix = z

@app.on_event("shutdown")
async def shutdown_event():
    await app.zabbix._aio_shutdown()

# List hosts and optionally their groups.
@api.get("/hosts/list")
async def api_hosts_list(group_ids:   Optional[List[int]] = Query(None),
                         list_groups: Optional[bool]      = False):
    data = await app.zabbix.host.get(
        groupids     = group_ids,
        output       = [ "name" ],
        selectGroups = [ "name" ] if list_groups else None
    )

    return data

# Get information about a specific host.
@api.get("/hosts/get")
async def api_hosts_get(host_id:     int,
                        list_groups: Optional[bool] = False):
    data = await app.zabbix.host.get(
        hostids      = [ host_id ],
        output       = "extend",
        selectGroups = [ "name" ] if list_groups else None
    )

    return data

# List triggers for a specific host.
@api.get("/items/list")
async def api_items_list(host_id: int):
    data = await app.zabbix.item.get(
        hostids = [ host_id ],
        output  = "extend",
    )

    return data

# List triggers for a specific host.
@api.get("/items/list")
async def api_items_list(host_id: int,
                         search:  Optional[str] = None):
    data = await app.zabbix.item.get(
        hostids = [ host_id ],
        output  = "extend",
    )

    return data

# Expand a specific trigger expression for a specific host.
@api.get("/macro/expand")
async def api_items_expand(host_id:    int,
                           expression: Optional[List[str]] = Query(None)):
    return {
        "status":  "failure",
        "message": "not implemented",
        "test_0":  host_id,
        "test_1":  expression
    }

# Serve the API.
app.include_router(
    router = api,
    prefix = "/api/v1"
)

# Serve the frontend.
app.mount(
    path = "/",
    name = "frontend",
    app  = StaticFiles(directory = "static", html = True)
)
