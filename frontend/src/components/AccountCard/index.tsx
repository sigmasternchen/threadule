import {FunctionComponent} from "react";
import {Avatar, Card, CardActions, CardContent, CardHeader, IconButton} from "@material-ui/core";
import Account from "../../api/entities/Account";
import ThreadList from "../ThreadList";

import AddIcon from '@material-ui/icons/Add';

export type AccountCardProps = {
    account: Account
}

const AccountCard: FunctionComponent<AccountCardProps> = ({account}) => {
    return (
        <Card>
            <CardHeader
                avatar={
                    <Avatar alt={account.screen_name} src={account.avatar_url}/>
                }
                action={
                    <IconButton aria-label="settings">
                    </IconButton>
                }
                title={account.name}
                subheader={account.screen_name}
            />
            <CardContent>
                <ThreadList threads={account.threads} />
            </CardContent>
            <CardActions disableSpacing>
                <IconButton aria-label="add">
                    <AddIcon/>
                </IconButton>
            </CardActions>
        </Card>
    )
}

export default AccountCard