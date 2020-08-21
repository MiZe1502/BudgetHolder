import { css } from 'emotion'

export { FlexVert } from "../../../common/styles/positioning";
import colors from "../../../common/styles/colors";

export const Form = css`
    width: 500px;

    padding: 16px;
    
    margin: 150px auto auto;
    border: 1px solid ${colors.greenPrimary};
    box-sizing: border-box;
    border-radius: 4px;
    box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
`;

export const ButtonsBlock = css`
    display: flex;
    justify-content: flex-end;
    margin-top: 8px;
`;

export const AuthButton = css`
    width: 100px;
    height: 40px;
`;