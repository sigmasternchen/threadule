import {FunctionComponent, useEffect, useState} from "react";
import {CircularProgress, Grid, Typography} from "@material-ui/core";
import Account from "../../api/entities/Account";
import {useAuth} from "../../auth/AuthProvider";
import AccountEndpoint from "../../api/endpoints/AccountEndpoint";
import {ClosedMessageBox, MessageBox, MessageBoxProps} from "../../components/MessageBox";
import AccountCard from "../../components/AccountCard";

type DashboardProps = {}

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
                console.dir(accounts)
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

    return <Grid container style={{
        marginTop: "40px",
        boxSizing: "border-box",
    }}>
        {loading &&
        <CircularProgress/>
        }
        {!loading && accounts.length == 0 &&
        <Typography variant={"h3"} style={{
            color: "grey",
        }}>
            No accounts yet.
        </Typography>
        }
        <Grid item container spacing={4}>
            {
                accounts.map((account) => (
                    <Grid item xs={4} key={account.id}>
                        <AccountCard account={account}/>
                    </Grid>
                ))
            }
        </Grid>

        <MessageBox {...message} />
    </Grid>
}

export default Dashboard