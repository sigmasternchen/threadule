import {FunctionComponent} from "react";
import {Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@material-ui/core";

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
    return (
        <Dialog open={open}>
            <DialogTitle title={"New Account"}/>
            <DialogContent>

            </DialogContent>
            <DialogActions>
                <Button onClick={() => {
                    onCancel()
                }} color="secondary">
                    Cancel
                </Button>
                <Button onClick={() => {
                   onSuccess()
                }} color="primary">
                    Submit
                </Button>
            </DialogActions>
        </Dialog>
    )
}

export default AccountDialog