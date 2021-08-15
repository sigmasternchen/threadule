import {Client} from "../client";
import {AxiosRequestConfig} from "axios";
import ErrorStatus, {ErrorStatusGuard} from "../entities/ErrorStatus";

abstract class Endpoint {
    protected client: Client

    protected constructor(client: Client) {
        this.client = client
    }

    protected requireAuthenticated() {
        if (!this.client.authorized)
            throw "this endpoint needs an authorized client"
    }

    private static handleResponse<T>(response: any): T {
        if (ErrorStatusGuard(response)) {
            throw (response as ErrorStatus)
        }

        return response as T
    }

    protected async get<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.get(path, config)
        )
    }

    protected async post<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.post(path, config)
        )
    }

    protected async options<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.options(path, config)
        )
    }

    protected async head<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.head(path, config)
        )
    }

    protected async put<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.put(path, config)
        )
    }

    protected async delete<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            this.client.axios.delete(path, config)
        )
    }
}

export default Endpoint