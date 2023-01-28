import { MantineProvider } from "@mantine/core";
import React from "react";
import { StoresContext } from "../Logic/provider";
import AppStore from "../Logic/store";

interface ProvidedAppProps {
  children?: React.ReactNode;
}

function ProvidedApp(props: ProvidedAppProps) {
  return (
    <StoresContext.Provider value={{
      appStore: new AppStore()
    }}>
      <MantineProvider withGlobalStyles withNormalizeCSS>
        {props.children}
      </MantineProvider>
    </StoresContext.Provider>
  );
}

export default ProvidedApp;
