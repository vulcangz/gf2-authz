import request from "utils/request";
import { baseUrl } from "service/common/api"
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { Principal } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";
import { PrincipalFormData } from "form/principal";


export const getPrincipals = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest,
) => {
    return paginate<Paginated<Principal>>({
        url: baseUrl() + '/principals',
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
}

export const createPrincipal = (
    token: string,
    data: PrincipalFormData,
) => {
    const url = new URL(baseUrl() + '/principals');

    return request.post<Principal | APIError>(url.href, data);
}

export const updatePrincipal =  (
    token: string,
    identifier: string,
    data: PrincipalFormData,
) => {
    const url = new URL(baseUrl() + `/principals/${identifier}`);

    return request.put<Principal>(url.href, data);
}

export const getPrincipal = (
    token: string,
    identifier: string,
) => {
    const url = new URL(baseUrl() + `/principals/${identifier}`);

    return request.get<Principal>(url.href);
}


export const deletePrincipal = (token: string, identifier: string) => {
    const url = new URL(baseUrl() + `/principals/${identifier}`);

    return request.delete<Principal>(url.href)

}
