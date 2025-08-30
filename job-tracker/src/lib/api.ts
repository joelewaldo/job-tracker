const BASE_URL = import.meta.env.VITE_API_URL;

async function request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const res = await fetch(`${BASE_URL}${endpoint}`, {
        headers: {
            "Content-Type": "application/json",
            ...(options?.headers || {}),
        },
        ...options,
    });

    if (!res.ok) {
        const message = await res.text();
        throw new Error(message || `HTTP error ${res.status}`);
    }

    return res.json();
}

export const api = {
    get: <T>(endpoint: string) => request<T>(endpoint),
    post: <T>(endpoint: string, body: any) =>
        request<T>(endpoint, { method: "POST", body: JSON.stringify(body) }),
    patch: <T>(endpoint: string, body: any) =>
        request<T>(endpoint, { method: "PATCH", body: JSON.stringify(body) }),
    delete: <T>(endpoint: string) =>
        request<T>(endpoint, { method: "DELETE" }),
};
