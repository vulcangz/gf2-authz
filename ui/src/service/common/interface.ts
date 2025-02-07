export interface ListResult<T = any> {
    total: number;
    page: number;
    size: number;
    list: T[];
}
