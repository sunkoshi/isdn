import { AppShell, Header } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import { useStores } from "./Logic/provider";
import NavigationBar from "./Organs/NavigationBar";
import ListLCIndex from "./Pages/ListSL/Index";
import ServerlessIndex from "./Pages/ServerlessCreate/Index";

export default function App() {
  const { appStore } = useStores();
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar />}
      header={
        <Header height={60} p="xs">
          <b>ISDN User Interface</b>
        </Header>
      }
      styles={(theme) => ({
        main: {
          backgroundColor:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[1]
        }
      })}
    >
      <Observer>
        {() => {
          if (appStore.navigationState === 0) return <ServerlessIndex />;
          return <ListLCIndex />;
        }}
      </Observer>
    </AppShell>
  );
}
