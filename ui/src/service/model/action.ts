import request from "utils/request";
import { baseUrl } from "service/common/api";
import { paginate, Paginated } from "service/common/paginate";
import { Action } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";

export const getActions = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest,
) => {
    const url = new URL(baseUrl() + "/actions");
    
    return paginate<Paginated<Action>>({
        url: baseUrl() + '/actions',
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
};

export const getAction = (identifier: string ) => {
    const url = new URL(baseUrl() + `/actions/${identifier}`);
    return request.post<Paginated<Action>>(url.href);
};

export const deleteAction = (identifier: string ) => {
    const url = new URL(baseUrl() + `/actions/${identifier}`);
    request.delete<Action>(url.href);
};

