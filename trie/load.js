import { check } from "k6";
import http from "k6/http";

export default function() {
    let res = http.get("https://2x7bqep62m.execute-api.eu-west-1.amazonaws.com/prefix");
    check(res, {
        "is status 200": (r) => r.status === 200
    });
};