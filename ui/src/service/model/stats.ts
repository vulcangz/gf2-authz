import useSWR from "swr";
import request from "utils/request";
import { baseUrl } from "service/common/api";
import type * as Type from 'service/common/interface';
import { StatsListResult } from "./model";

export const getStats = () => {
    const url = new URL(baseUrl() + "/stats");
    return request.get<StatsListResult>(url.href);
};

export const useStatList = () => {
    const url = new URL(baseUrl() + "/stats");
    const { data, error } = useSWR<StatsListResult, Error>(url.href, (url) =>
        request.get(url.href, { allow404: true })
    );
    return {
        data,
        isLoading: !data && !error,
        error,
    };
};