import request from "utils/request";
import { baseUrl } from "service/common/api";
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { Client } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";
import { ClientFormData } from "form/client";

export const getClients = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest
) => {
    return paginate<Client>({
        url: baseUrl() + "/clients",
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
};

export const createClient = (
    token: string,
    data: ClientFormData
) => {    
    const url = new URL(baseUrl() + "/clients");

    return request.post<Client | APIError>(url.href, data)
}

export const getClient = (
    token: string,
    identifier: string
) => {
    const url = new URL(baseUrl() + `/clients/${identifier}`);
    return request.get<Client>(url.href);
}

export const deleteClient = (
    token: string,
    identifier: string
) => {
    const url = new URL(baseUrl() + `/clients/${identifier}`);
    return request.delete<Client>(url.href)    
}
