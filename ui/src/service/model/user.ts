import request from "utils/request";
import { baseUrl } from "service/common/api"
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { User } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";
import { UserFormData } from "form/user";


export const getUsers = async (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest,
) => {
    return paginate<User>({
        url: baseUrl() + '/users',
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
}

export const createUser = (
    token: string,
    data: UserFormData,
) => {
    const url = new URL(baseUrl() + '/users');

    return request.post<User | APIError>(url.href, data)

}

export const getUser = (
    token: string,
    identifier: string,
) => {
    const url = new URL(baseUrl() + `/users/${identifier}`);

    return request.get<User>(url.href)

}

export const deleteUser = async (token: string, identifier: string) => {
    const url = new URL(baseUrl() + `/users/${identifier}`);

    return request.delete<User>(url.href)
}
