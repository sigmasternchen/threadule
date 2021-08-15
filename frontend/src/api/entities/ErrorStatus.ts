type ErrorStatus = {
    status: number,
    message: string,
    details: string,
    time: Date
}

export const ErrorStatusGuard = (value: any): value is ErrorStatus => {
    return value.status !== undefined &&
        typeof value.status == "number" &&
        value.message !== undefined &&
        value.details !== undefined &&
        value.time !== undefined;
}


export default ErrorStatus