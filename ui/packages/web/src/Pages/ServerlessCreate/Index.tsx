import {
  Button,
  FileButton,
  Group,
  NumberInput,
  Select,
  TextInput,
  Text,
  Divider,
} from "@mantine/core";
import React, { useState } from "react";
import CodeEditor from "./CodeEditor";

function ServerlessIndex() {
  const [file, setFile] = useState<File | null>(null);
  const [runtime, setRuntime] = useState("js");
  return (
    <div>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <TextInput
          style={{ width: "32%" }}
          placeholder="Name"
          label="Name"
          withAsterisk
        />
        <Select
          withAsterisk
          defaultValue={runtime}
          onChange={(e) => {
            setRuntime(e || "js");
          }}
          style={{ width: "32%" }}
          label="Select Runtime"
          placeholder="Pick one"
          data={[
            { value: "js", label: "NodeJS 14" },
            { value: "py", label: "Python 3" },
            { value: "cpp", label: "C++ 11" },
          ]}
        />
        <NumberInput
          style={{ width: "32%" }}
          defaultValue={60}
          placeholder="Enter timeout"
          label="Timeout (in seconds)"
          withAsterisk
        />
      </div>
      <CodeEditor language={runtime} />
      <Divider style={{ margin: "10px 0" }} />
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <Group position="center">
          <FileButton onChange={setFile}>
            {(props) => (
              <Button variant="default" {...props}>
                Upload Code / Zip
              </Button>
            )}
          </FileButton>
        </Group>
        {file && (
          <Text size="sm" align="center" mt="sm">
            Picked file: {file.name}
          </Text>
        )}
        <Button>Create Serverless Load</Button>
      </div>
    </div>
  );
}

export default ServerlessIndex;
