import {FunctionComponent, useEffect, useState} from "react";
import {
    Button,
    Dialog,
    DialogContent,
    DialogTitle,
    Step,
    StepContent,
    StepLabel,
    Stepper,
    TextField,
    Typography
} from "@material-ui/core";
import {useAuth} from "../../auth/AuthProvider";
import AccountEndpoint from "../../api/endpoints/AccountEndpoint";
import {MessageBox} from "../MessageBox";

export type AccountFormDialogProps = {
    open: boolean
    onSuccess: () => void
    onCancel: () => void
}

const AccountDialog: FunctionComponent<AccountFormDialogProps> = (
    {
        open,
        onSuccess,
        onCancel
    }
) => {
    const {client} = useAuth()

    const [activeStep, setActiveStep] = useState<number>(0)

    const [accountId, setAccountId] = useState<string | null>(null)
    const [code, setCode] = useState<string>("")

    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        if (!open) {
            // dialog is closed; reset component
            setError(null)
            setCode("")
            setAccountId(null)
            setActiveStep(0)
        }
    })

    return (
        <Dialog open={open}>
            <DialogTitle title={"New Account"}/>
            <DialogContent>
                <Stepper activeStep={activeStep} orientation="vertical">
                    <Step>
                        <StepLabel>Begin</StepLabel>
                        <StepContent>
                            <Typography>
                                This dialog will guide you through the process of adding a new account.
                            </Typography>
                            <Button
                                onClick={() => {
                                    onCancel()
                                }}
                                color="secondary"
                            >
                                Cancel
                            </Button>
                            <Button
                                onClick={() => {
                                    setActiveStep(activeStep + 1)
                                }}
                                color="primary"
                            >
                                Next
                            </Button>
                        </StepContent>
                    </Step>
                    <Step>
                        <StepLabel>Authorize The App</StepLabel>
                        <StepContent>
                            <Typography>
                                When you click "Next" you will be sent to Twitter.<br/>
                                After you authorize our app you will be given a 6-digit authorization code that you'll
                                need in the next step.
                            </Typography>
                            <Button
                                onClick={() => {
                                    setActiveStep(activeStep - 1)
                                }}
                                color="secondary"
                            >
                                Back
                            </Button>
                            <Button
                                onClick={() => {
                                    new AccountEndpoint(client).addAccount()
                                        .then((data) => {
                                            setAccountId(data.id)
                                            setActiveStep(activeStep + 1)
                                            window.open(data.url, "_blank")
                                        })
                                        .catch(error => {
                                            console.error(error)
                                            setError("Something went wrong.")
                                        })
                                }}
                                color="primary"
                            >
                                Next
                            </Button>
                        </StepContent>
                    </Step>
                    <Step>
                        <StepLabel>Enter Authorization Code</StepLabel>
                        <StepContent>
                            <Typography>
                                Please enter the authorization code that twitter gave you.
                            </Typography>
                            <TextField
                                fullWidth
                                label="Authorization Code"
                                variant="outlined"
                                value={code}
                                error={!(code.length == 0 || code.match(/^[0-9]{6,}$/))}
                                onChange={(event) => {
                                    setCode(event.target.value)
                                }}
                            />
                            <Button
                                onClick={() => {
                                    setActiveStep(activeStep - 1)
                                }}
                                color="secondary"
                            >
                                Back
                            </Button>
                            <Button
                                disabled={!code.match(/^[0-9]{6,}$/)}
                                onClick={() => {
                                    new AccountEndpoint(client).addAccountResolve(accountId!, code)
                                        .then(() => {
                                            onSuccess()
                                        })
                                        .catch(error => {
                                            console.error(error)
                                            setError("Something went wrong.")
                                        })
                                }}
                                color="primary"
                            >
                                Done
                            </Button>
                        </StepContent>
                    </Step>
                </Stepper>
            </DialogContent>
            <MessageBox open={Boolean(error)} success={false} message={error || ""} onClose={() => {
                setError(null)
            }}/>
        </Dialog>
    )
}

export default AccountDialog