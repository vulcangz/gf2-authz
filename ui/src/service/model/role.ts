import request from "utils/request";
import { RoleFormData } from "../../form/role";
import { baseUrl } from "service/common/api"
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { Role } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";

export const getRoles = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest,
) => {
    return paginate<Paginated<Role>>({
        url: baseUrl() + '/roles',
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
}

export const getRole = (
    token: string,
    identifier: string,
) => {
    const url = new URL(baseUrl() + `/roles/${identifier}`);

    return request.get<Role>(url.href);

}

export const createRole = (
    token: string,
    data: RoleFormData,
) => {
    const url = new URL(baseUrl() + '/roles');

    return request.post<Role | APIError>(url.href, data);
}

export const updateRole = (
    token: string,
    identifier: string,
    data: RoleFormData,
) => {
    const url = new URL(baseUrl() + `/roles/${identifier}`);

    return request.put<Role>(url.href, data);
}

export const deleteRole = (token: string, identifier: string) => {
    const url = new URL(baseUrl() + `/roles/${identifier}`);

    return request.delete<Role>(url.href);
}
