import request from "utils/request";
import { SigninFormData } from "form/signin";
import { baseUrl } from "service/common/api";
import { AuthResponse } from "service/auth/model";

export const signIn = (params: SigninFormData) => {
  const url = new URL(baseUrl() + "/auth");
  return request.post<AuthResponse>(url.href, params);
};
