import { css } from 'emotion'

export { Font312DarkGray, Font316Black } from "../../../styles/fonts";
import colors from "../../../styles/colors";

export const Input = css`
    height: 30px;
    border-radius: 4px;
    padding: 4px 8px;
    border: 1px solid ${colors.greenPrimary};
    margin-bottom: 8px;
    caret-color: ${colors.darkGray};
    outline: none;
    text-decoration: none;
    
    &:focus {
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.15);
        border: 1px solid ${colors.greenActive};
    }
    
    &:disabled {
        border: 1px solid ${colors.darkGray};
        background: ${colors.gray};
    }
`