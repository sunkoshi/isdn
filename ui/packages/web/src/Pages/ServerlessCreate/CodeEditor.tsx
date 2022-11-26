import React from "react";
import TextAreaCodeEditor from "@uiw/react-textarea-code-editor";

interface CodeEditorProps {
  language?: string;
}

function CodeEditor(props: CodeEditorProps) {
  const [code, setCode] = React.useState(
    `function add(a, b) {\n  return a + b;\n}`
  );
  return (
    <div>
      <TextAreaCodeEditor
        value={code}
        language={props.language}
        placeholder={`Write your .${props.language} code here.`}
        onChange={(evn) => setCode(evn.target.value)}
        padding={15}
        style={{
          borderRadius: "5px",
          border: "1px solid #c7c7c7",
          fontSize: 16,
          height: "40vh",
          marginTop: "20px",
          backgroundColor: "#e2e2e2",
          fontFamily:
            "ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace",
        }}
      />
    </div>
  );
}

export default CodeEditor;
