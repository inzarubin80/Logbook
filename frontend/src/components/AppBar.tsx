import * as React from "react";
import { styled, useTheme } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import CssBaseline from "@mui/material/CssBaseline";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import InboxIcon from "@mui/icons-material/MoveToInbox";
import DraftsIcon from "@mui/icons-material/Drafts";
import { Routes, Route } from "react-router-dom";
import Сategories from "./Сategories";
import SportSchools from "./SportSchools";
import TypeTournaments from "./TypeTournaments";
import Coaches from "./Coaches";
import CategoryValue from "./CategoryValue";
import ScoreScale from "./ScoreScale";
import Sportsman from './Sportsman';
import Tournament from "./Tournament";


import { Anon, LoggedIn } from "../Shared/Roles";
import LogIn from "./login/login";
import Verify from "./Verify/Verify";

import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from "react-router-dom";

const drawerWidth = 250;

const Main = styled("main", { shouldForwardProp: (prop) => prop !== "open" })<{
  open?: boolean;
}>(({ theme, open }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  transition: theme.transitions.create("margin", {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  marginLeft: `-${drawerWidth}px`,
  ...(open && {
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  }),
}));

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== "open",
})<AppBarProps>(({ theme, open }) => ({
  transition: theme.transitions.create(["margin", "width"], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: `${drawerWidth}px`,
    transition: theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

interface ListItemLinkProps {
  icon?: React.ReactElement;
  primary: string;
  to: string;
}

function ListRouter() {
  return (
    <List aria-label="main mailbox folders">
      <ListItemLink to="/category" primary="Категории" icon={<InboxIcon />} />
      <ListItemLink to="/sportSchool" primary="Школы" icon={<InboxIcon />} />
      <ListItemLink to="/coache" primary="Тренеры" icon={<InboxIcon />} />
      <ListItemLink
        to="/typeTournament"
        primary="Типы турниров"
        icon={<InboxIcon />}
      />
      <ListItemLink to="/login" primary="Вход" icon={<DraftsIcon />} />
      <ListItemLink
        to="/сategoryValue"
        primary="Значение категорий"
        icon={<DraftsIcon />}
      />
       <ListItemLink
        to="/scoreScale"
        primary="Шкала оценок"
        icon={<DraftsIcon />}
      />
       <ListItemLink
        to="/sportsman"
        primary="Спортсмены"
        icon={<DraftsIcon />}
      />
       <ListItemLink
        to="/tournament"
        primary="Турниры"
        icon={<DraftsIcon />}
      />
    </List>
  );
}

const Link = React.forwardRef<HTMLAnchorElement, RouterLinkProps>(function Link(
  itemProps,
  ref
) {
  return <RouterLink ref={ref} {...itemProps} role={undefined} />;
});

function ListItemLink(props: ListItemLinkProps) {
  const { icon, primary, to } = props;

  return (
    <li>
      <ListItem button component={Link} to={to}>
        {icon ? <ListItemIcon>{icon}</ListItemIcon> : null}
        <ListItemText primary={primary} />
      </ListItem>
    </li>
  );
}

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

export default function MyBar() {
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <Box sx={{ display: "flex" }}>
      <CssBaseline />
      <AppBar position="fixed" open={open}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{ mr: 2, ...(open && { display: "none" }) }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div">
            Рейтинг спортсменов
          </Typography>
        </Toolbar>
      </AppBar>

      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: drawerWidth,
            boxSizing: "border-box",
          },
        }}
        variant="persistent"
        anchor="left"
        open={open}
      >
        <DrawerHeader>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === "ltr" ? (
              <ChevronLeftIcon />
            ) : (
              <ChevronRightIcon />
            )}
          </IconButton>
        </DrawerHeader>
        <Divider />

        <ListRouter />
      </Drawer>

      <Main open={open}>
        <DrawerHeader />

        <Routes>
          <Route path="/" element={<Сategories />} />
          <Route path="/category" element={<Сategories />} />
          <Route path="/sportSchool" element={<SportSchools />} />
          <Route path="/typeTournament" element={<TypeTournaments />} />
          <Route path="/coache" element={<Coaches />} />
          <Route path="/login" element={<LogIn />} />
          <Route path="/verify/:code" element={<Verify />} />
          <Route path="/сategoryValue" element={<CategoryValue />} />
          <Route path="/scoreScale" element={<ScoreScale/>} />
          <Route path="/sportsman" element={<Sportsman/>} />
          <Route path="/tournament" element={<Tournament/>} />
        </Routes>
      </Main>
    </Box>
  );
}
