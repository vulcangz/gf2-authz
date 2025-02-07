import { baseUrl } from "service/common/api";
import { FilterRequest } from "service/common/filter";
import { paginate, Paginated } from "service/common/paginate";
import { SortRequest } from "service/common/sort";
import { Audit } from "service/model/model";

export const getAudits = async (
    token: string,
    page?: number,
    size?: number,
    filter?: FilterRequest,
    sort?: SortRequest
) => {
    return paginate<Paginated<Audit>>({
        url: baseUrl() + "/audits",
        token: token,
        page: page,
        size: size,
        filter: filter,
        sort: sort,
    });
};
