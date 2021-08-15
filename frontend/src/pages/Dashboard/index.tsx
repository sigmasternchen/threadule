import {FunctionComponent, useEffect, useState} from "react";
import {Avatar, Card, CardHeader, CircularProgress, Grid, IconButton, Typography} from "@material-ui/core";
import Account from "../../api/entities/Account";
import {useAuth} from "../../auth/AuthProvider";
import AccountEndpoint from "../../api/endpoints/AccountEndpoint";
import {ClosedMessageBox, MessageBox, MessageBoxProps} from "../../components/MessageBox";

type DashboardProps = {

}

const Dashboard: FunctionComponent<DashboardProps> = () => {
    const [loading, setLoading] = useState<boolean>(true)
    const [accounts, setAccounts] = useState<Account[]>([])

    const [message, setMessage] = useState<MessageBoxProps>({
        ...ClosedMessageBox,
        onClose: () => {
            setMessage({
                ...message,
                open: false
            })
        }
    })

    const {client} = useAuth()
    const accountEndpoint = new AccountEndpoint(client)

    useEffect(() => {
        accountEndpoint.getAll()
            .then(accounts => {
                setAccounts(accounts)
                setLoading(false)
            })
            .catch(reason => {
                console.error(reason)
                setLoading(false)
                setMessage({
                    ...message,
                    open: true,
                    success: false,
                    message: "Couldn't fetch accounts."
                })
            })
    }, [])

    return <div style={{
        marginTop: "40px"
    }}>
        { loading &&
            <CircularProgress />
        }
        { !loading && accounts.length == 0 &&
            <Typography variant={"h3"} style={{
                color: "grey",
            }}>
                No accounts yet.
            </Typography>
        }
        <Grid container xs={12} spacing={4}>
        {
            accounts.map((account) => {
                return <>
                    <Grid item xs={4}>
                        <Card>
                            <CardHeader
                                avatar={
                                    <Avatar alt={account.screen_name} src={account.avatar_url} />
                                }
                                action={
                                    <IconButton aria-label="settings">
                                    </IconButton>
                                }
                                title={account.name}
                                subheader={account.screen_name}
                            />
                        </Card>
                    </Grid>
                </>
            })
        }
        </Grid>

        <MessageBox {...message} />
    </div>
}

export default Dashboard