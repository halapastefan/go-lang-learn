import { check } from "k6";
import http from "k6/http";

export default function() {
    let res = http.get("http://localhost:3300/call/939898123985/2019-12-04T00:00:00.00Z");
    check(res, {
        "is status 200": (r) => r.status === 200
    });
};