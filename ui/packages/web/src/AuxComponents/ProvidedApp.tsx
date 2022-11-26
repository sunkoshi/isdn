import { MantineProvider } from "@mantine/core";
import React from "react";

interface ProvidedAppProps {
  children?: React.ReactNode;
}

function ProvidedApp(props: ProvidedAppProps) {
  return (
    <>
      <MantineProvider withGlobalStyles withNormalizeCSS>
        {props.children}
      </MantineProvider>
    </>
  );
}

export default ProvidedApp;
