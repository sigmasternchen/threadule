import {FunctionComponent} from "react";
import {Card, CardContent, CardHeader, Grid} from "@material-ui/core";

export type SettingsCardProps = {
    name: string
}

const SettingsCard: FunctionComponent<SettingsCardProps> = ({name, children}) => {
    return (
        <Grid item xs={6}>
            <Card>
                <CardHeader
                    title={name}
                />
                <CardContent>
                    {children}
                </CardContent>
            </Card>
        </Grid>
    )
}

export default SettingsCard