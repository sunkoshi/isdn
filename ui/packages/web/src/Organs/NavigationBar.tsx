import { Navbar, NavLink } from "@mantine/core";
import { IconList, IconPlus } from "@tabler/icons";
import { Observer } from "mobx-react-lite";
import { useStores } from "../Logic/provider";
function NavigationBar() {
  const { appStore } = useStores();
  return (
    <Navbar width={{ base: 300 }} height={"100%"} p="xs">
      <Observer>
        {() => {
          const { navigationState } = appStore;
          return (
            <>
              <NavLink
                active={navigationState === 0}
                onClick={() => appStore.setNavigationState(0)}
                label="Create SL"
                icon={<IconPlus size={16} stroke={1.5} />}
              />
              <NavLink
                active={navigationState === 1}
                onClick={() => appStore.setNavigationState(1)}
                label="List SL"
                icon={<IconList size={16} stroke={1.5} />}
              />
            </>
          );
        }}
      </Observer>
    </Navbar>
  );
}

export default NavigationBar;
