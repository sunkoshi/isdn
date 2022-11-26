import { Navbar, NavLink } from "@mantine/core";
import { IconList, IconPlus } from "@tabler/icons";
function NavigationBar() {
  return (
    <Navbar width={{ base: 300 }} height={"100%"} p="xs">
      <NavLink
        label="Create SL"
        icon={<IconPlus size={16} stroke={1.5} />}
        variant="filled"
        active
      />
      <NavLink label="List SL" icon={<IconList size={16} stroke={1.5} />} />
    </Navbar>
  );
}

export default NavigationBar;
