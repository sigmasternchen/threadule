import {FunctionComponent} from "react";
import SettingsCard from "../../components/SettingsCard";
import UserSettings from "./UserSettings";
import {Grid} from "@material-ui/core";
import styles from "./UserSettings.module.css"

export type SettingsProps = {}

const Settings: FunctionComponent<SettingsProps> = () => {
    return (
        <>
            <Grid container className={styles.container} spacing={4}>
                <SettingsCard name={"User"}>
                    <UserSettings/>
                </SettingsCard>
            </Grid>
        </>
    )
}

export default Settings