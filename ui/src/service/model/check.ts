import request from "utils/request";
import { baseUrl } from "service/common/api";
import { CheckFormData } from "form/check";

export type Check = {
    is_allowed: boolean;
};

export type CheckResponse = {
    checks: Check[];
};

export const check = (token: string, data: CheckFormData) => {
    const url = new URL(baseUrl() + "/check");
    return request.post<CheckResponse>(url.href, data);
};
