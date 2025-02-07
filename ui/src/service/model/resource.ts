import request from "utils/request";
import { baseUrl } from "service/common/api"
import { paginate, Paginated } from "service/common/paginate";
import { APIError } from "service/error/model";
import { Resource } from "./model";
import { FilterRequest } from "service/common/filter";
import { SortRequest } from "service/common/sort";
import { ResourceFormData } from "form/resource";



export const getResources = (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest,
) => {
    return paginate<Paginated<Resource>>({
        url: baseUrl() + '/resources',
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
}

export const createResource = (
    token: string,
    data: ResourceFormData,
) => {
    const url = new URL(baseUrl() + '/resources');

    return request.post<Resource | APIError>(url.href, data);
}

export const updateResource = (
    token: string,
    identifier: string,
    data: ResourceFormData,
) => {
    const url = new URL(baseUrl() + `/resources/${identifier}`);

    return request.put<Resource>(url.href, data);

}

export const getResource = async (
    token: string,
    identifier: string,
) => {
    const url = new URL(baseUrl() + `/resources/${identifier}`);

    return request.get<Resource>(url.href);
}

export const deleteResource =  (token: string, identifier: string) => {    
    const url = new URL(baseUrl() + `/resources/${identifier}`);

    return request.delete<Resource>(url.href);
}
