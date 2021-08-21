class CustomDate extends Date {
    constructor(date?: Date) {
        super(date ? date : new Date())
    }

    public toLocalISOString(milliseconds: boolean = true,
                            timezone: boolean = true): string {
        let result =
            this.getFullYear() + "-" +
            (String(this.getMonth()).padStart(2, "0")) + "-" +
            (String(this.getDay()).padStart(2, "0")) + "T" +
            (String(this.getHours()).padStart(2, "0")) + ":" +
            (String(this.getMinutes()).padStart(2, "0")) + ":" +
            (String(this.getSeconds()).padStart(2, "0"))
        if (milliseconds) {
            result += "." + (String(this.getMilliseconds()).padStart(3, "0"))
        }
        if (timezone) {
            const offset = -this.getTimezoneOffset()
            const hourOffset = Math.floor(Math.abs(offset) / 60)
            const minuteOffset = offset % 60
            result +=
                (offset > 0 ? "+" : "-") +
                (String(hourOffset).padStart(2, "0")) +
                (String(minuteOffset).padStart(2, "0"))
        }

        return result
    }
}

export default CustomDate