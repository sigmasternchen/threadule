import {Button, Grid, LinearProgress} from "@material-ui/core";
import Alert from '@material-ui/lab/Alert';
import {Field, Formik, FormikHelpers} from "formik"
import {TextField} from "formik-material-ui";
import {useAuth} from "../../auth/AuthProvider";
import * as yup from "yup"
import {useState} from "react";
import {useHistory} from "react-router-dom";

type LoginFormProps = {
    username: string
    password: string
}

const LoginFormValidationSchema = yup.object({
    username: yup.string().required(),
    password: yup.string().required()
})

const Login = () => {

    const auth = useAuth()
    const history = useHistory()

    const [error, setError] = useState<string|null>(null)

    const onSubmit = async (values: LoginFormProps, helper: FormikHelpers<LoginFormProps>) => {
        try {
            const user = await auth.login(values.username, values.password)
            console.log(user)
            history.push("/");
        } catch (e) {
            helper.setSubmitting(false)
            setError("Login failed!")
            console.error(e)
        }
    }

    return (
        <Grid container justifyContent="center">
            <Grid item container xs={4} spacing={2}>
                <Formik<LoginFormProps>
                    initialValues={{
                        username: "",
                        password: ""
                    }}
                    onSubmit={onSubmit}
                    validationSchema={LoginFormValidationSchema}
                >
                    {({submitForm, isSubmitting}) => <>
                        <Grid item xs={12}>
                            <Field
                                component={TextField}
                                label="Username"
                                name="username"
                                variant="outlined"
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <Field
                                component={TextField}
                                label="Password"
                                name="password"
                                type="password"
                                variant="outlined"
                                fullWidth
                            />
                        </Grid>
                        { isSubmitting && <LinearProgress  /> }
                        <Grid item xs={12}>
                            <Button
                                variant={"contained"}
                                onClick={submitForm}
                                fullWidth
                            >
                                Submit
                            </Button>
                        </Grid>
                        { error && <Grid item xs={12}>
                            <Alert severity="error">{error}</Alert>
                        </Grid>}
                    </>}
                </Formik>
            </Grid>
        </Grid>
    )
}

export default Login