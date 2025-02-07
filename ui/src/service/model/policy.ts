import request from "utils/request";
import { baseUrl } from "service/common/api";
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { Policy } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";
import { PolicyFormData } from "form/policy";

export const getPolicies = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest
) => {
    return paginate<Paginated<Policy>>({
        url: baseUrl() + "/policies",
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
};

export const getPolicy = (identifier: string) => {
    const url = new URL(baseUrl() + `/policies/${identifier}`);

    return request.get<Policy>(url.href);
};

export const createPolicy = (data: PolicyFormData) => {
    const url = new URL(baseUrl() + "/policies");

    return request.post<Policy | APIError>(url.href, data);
};

export const updatePolicy = (
    identifier: string,
    data: PolicyFormData
) => {
    const url = new URL(baseUrl() + `/policies/${identifier}`);

    return request.put<Policy>(url.href, data);
};

export const deletePolicy = (identifier: string) => {
    const url = new URL(baseUrl() + `/policies/${identifier}`);

    return request.delete<Policy>(url.href);
};
