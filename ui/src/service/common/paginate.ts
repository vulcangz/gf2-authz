import request from 'utils/request';
import { FilterRequest, filterRequestToValue } from './filter';
import { SortRequest, sortRequestToValue } from './sort';

export type PaginateRequest = {
    url: string
    token: string
    page?: number
    size?: number
    filter?: FilterRequest    
    sort?: SortRequest
}

export type Paginated<T> = {
    id: any;
    total: number
    page: number
    size: number
    data: T[]
};

export const paginate = <T>(req: PaginateRequest) => {
    const url = new URL(req.url);

    if (req.page !== undefined) {
        url.searchParams.set('page', req.page.toString());
    }

    if (req.size !== undefined) {
        url.searchParams.set('size', req.size.toString());
    }

    if (req.filter) {
        url.searchParams.set('filter', filterRequestToValue(req.filter));
    }

    if (req.sort) {
        url.searchParams.set('sort', sortRequestToValue(req.sort));
    }

    const response = request.get<Paginated<T>>(url.href, req);
    
    return response;
}
