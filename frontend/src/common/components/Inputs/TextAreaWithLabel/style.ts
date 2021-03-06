import {css} from 'emotion'

import {
    Font312DarkGray, Font312Black, Font312RedAttention
} from "../../../styles/fonts";
import {InvalidInput} from "../../../styles/inputs";
import colors from "../../../styles/colors";

export const style = {
    Font312DarkGray,
    Font312Black,
    Font312RedAttention,

    InvalidInput,

    TextArea: css`
        border-radius: 4px;
        padding: 4px 8px;
        border: 1px solid ${colors.greenPrimary};
        margin-bottom: 8px;
        caret-color: ${colors.darkGray};
        outline: none;
        text-decoration: none;
        resize: none;
        box-shadow: inset 0 3px 0 0 rgba(29, 29, 27, 0.05), inset 0 2px 0 0 rgba(29, 29, 27, 0.05);
        
        &:focus {
            box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.15);
            border: 1px solid ${colors.greenActive};
        }
        
        &:disabled {
            border: 1px solid ${colors.darkGray};
            background: ${colors.gray};
        }
    `
}