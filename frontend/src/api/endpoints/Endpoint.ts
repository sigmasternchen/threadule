import {Client} from "../client";
import {AxiosRequestConfig, AxiosResponse} from "axios";
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

    private static handleResponse<T>(response: AxiosResponse<T|ErrorStatus>): T {
        if (ErrorStatusGuard(response.data)) {
            throw (response.data as ErrorStatus)
        }

        return response.data as T
    }

    protected async get<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.get<T|ErrorStatus>(path, config)
        )
    }

    protected async post<S, T>(path: string, data: S, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.post<T|ErrorStatus>(path, data, config)
        )
    }

    protected async options<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.options<T|ErrorStatus>(path, config)
        )
    }

    protected async head<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.head<T|ErrorStatus>(path, config)
        )
    }

    protected async put<S, T>(path: string, data: S, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.put<T|ErrorStatus>(path, data, config)
        )
    }

    protected async delete<T>(path: string, config?: AxiosRequestConfig): Promise<T> {
        return Endpoint.handleResponse<T>(
            await this.client.axios.delete<T|ErrorStatus>(path, config)
        )
    }
}

export default Endpoint