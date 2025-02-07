export const baseUrl = (): string | undefined => import.meta.env.REACT_APP_API_BASE_URI || 'http://localhost:8080/v1';