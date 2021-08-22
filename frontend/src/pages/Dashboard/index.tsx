import {FunctionComponent, useEffect, useState} from "react";
import {CircularProgress, Fab, Grid, Typography} from "@material-ui/core";
import Account from "../../api/entities/Account";
import {useAuth} from "../../auth/AuthProvider";
import AccountEndpoint from "../../api/endpoints/AccountEndpoint";
import {ClosedMessageBox, MessageBox, MessageBoxProps} from "../../components/MessageBox";
import AccountCard from "../../components/AccountCard";
import AddIcon from "@material-ui/icons/Add";
import styles from "./Dashboard.module.css"
import AccountDialog from "../../components/AccountDialog";

type DashboardProps = {}

const Dashboard: FunctionComponent<DashboardProps> = () => {
    const [_reload, _setReload] = useState<number>(0)
    const reload = () => {
        _setReload(_reload + 1)
    }
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

    const [openNewAccountDialog, setOpenNewAccountDialog] = useState<boolean>(false)

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
    }, [_reload])

    return (
        <>
            <Grid container className={styles.container}>
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
                        accounts.map((account, index) => (
                            <Grid item xs={4} key={account.id}>
                                <AccountCard
                                    account={account}
                                    onUpdate={(account: Account) => {
                                        accounts[index] = account
                                        console.log(accounts)
                                        setAccounts([...accounts])
                                    }}
                                />
                            </Grid>
                        ))
                    }
                </Grid>
            </Grid>

            <AccountDialog
                open={openNewAccountDialog}
                onSuccess={() => {
                    setOpenNewAccountDialog(false)
                    reload()
                }}
                onCancel={() => {
                    setOpenNewAccountDialog(false)
                }}
            />

            <Fab
                size={"large"}
                color="primary"
                aria-label="add"
                className={styles.addButton}
                onClick={() => {
                    setOpenNewAccountDialog(true)
                }}
            >
                <AddIcon/>
            </Fab>

            <MessageBox {...message} />
        </>
    )
}

export default Dashboard