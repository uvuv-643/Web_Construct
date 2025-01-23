import {Button, Modal} from '@gravity-ui/uikit';
import {useState} from 'react';
import './ComponentCard.css';
import AceEditor from 'react-ace';

import 'ace-builds/src-noconflict/mode-text';
import 'ace-builds/src-noconflict/theme-github';
import 'ace-builds/src-noconflict/ext-language_tools';

export type ComponentDto = {
    id: string;
    request: string;
    response: string;
    modified_response: string;
    created_at: string;
    modified_at: string;
};

type ComponentCardProps = {
    component: ComponentDto;
};

export const ComponentCard = ({component}: ComponentCardProps) => {
    const [open, setOpen] = useState(false);
    const [openCode, setOpenCode] = useState(false);

    return (
        <li className="card">
            <h1>{component.id}</h1>
            <p>{component.request}</p>
            <div className="card-buttons">
                <Button view="action" onClick={() => setOpen(true)}>
                    Демо
                </Button>
                <Button view="action" onClick={() => setOpenCode(true)}>
                    Код
                </Button>
            </div>
            <Modal open={open} onClose={() => setOpen(false)}>
                <div
                    className="component-wrapper"
                    dangerouslySetInnerHTML={{__html: component.response}}
                ></div>
            </Modal>
            <Modal open={openCode} onClose={() => setOpenCode(false)}>
                <div className="component-wrapper">
                    <AceEditor
                        mode="text"
                        theme="github"
                        value={component.response}
                        width="700px"
                        onChange={(newValue) => {
                            console.log(newValue);
                        }}
                        name="UNIQUE_ID_OF_DIV"
                        editorProps={{$blockScrolling: true}}
                    />
                </div>
            </Modal>
        </li>
    );
};
