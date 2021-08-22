import {AppBar, IconButton, Menu, MenuItem, Toolbar, Typography} from "@material-ui/core";
import MenuIcon from '@material-ui/icons/Menu';
import {FunctionComponent, useState} from "react";
import {useAuth} from "../../auth/AuthProvider";
import {AccountCircle} from "@material-ui/icons";
import {Link, useHistory} from "react-router-dom";

type MenuBarProps = {
    pageName: string
}

const MenuBar: FunctionComponent<MenuBarProps> = ({pageName}) => {
    const { user, logout } = useAuth()
    const history = useHistory()

    const [anchorEl, setAnchorEl] = useState<HTMLElement|null>(null)
    const open = Boolean(anchorEl);

    const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <AppBar position="static">
            <Toolbar>
                <IconButton edge="start" color="inherit" aria-label="menu">
                    <MenuIcon />
                </IconButton>
                <Typography variant="h6" style={{flexGrow: 1}}>
                    <Link to={"/"} style={{
                        color: "inherit",
                        textDecoration: "inherit"
                    }}>
                        Threadule - {pageName}
                    </Link>
                </Typography>
                <IconButton
                    aria-label="account of current user"
                    aria-controls="menu-appbar"
                    aria-haspopup="true"
                    onClick={handleMenu}
                    color="inherit"
                >
                    <AccountCircle />
                </IconButton>
                <Menu
                    id="menu-appbar"
                    anchorEl={anchorEl}
                    anchorOrigin={{
                        vertical: 'top',
                        horizontal: 'right',
                    }}
                    keepMounted
                    transformOrigin={{
                        vertical: 'top',
                        horizontal: 'right',
                    }}
                    open={open}
                    onClose={handleClose}
                >
                    <MenuItem onClick={() => history.push("/settings")}>Settings</MenuItem>
                    <MenuItem onClick={logout}>Logout</MenuItem>
                </Menu>
            </Toolbar>
        </AppBar>
    )
}

export default MenuBar