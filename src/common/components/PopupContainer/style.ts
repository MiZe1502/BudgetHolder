import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font724Black, Font312Red } from "../../styles/fonts";
export { FlexHorCenter, Overflowed, FlexVert } from "../../styles/positioning";

export const Popup = css`
    position: fixed;
    margin: auto;
    width: 900px;
    height: auto;
    border-radius: 4px;
    background: ${colors.white};
    box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
`;

export const UtilBlock = css`
    width: 100%;
    height: 50px;
    background-color: ${colors.gray};
    padding: 16px;
    box-sizing: border-box;
`;

export const Header = css`
    cursor: grab;
    border-top-right-radius: 4px;
    border-top-left-radius: 4px;
    justify-content: space-between;
`;

export const Footer = css`
    border-bottom-right-radius: 4px;
    border-bottom-left-radius: 4px;
    height: auto;
    justify-content: flex-end;
`;

export const FooterWithErrors = css`
    justify-content: space-between;
`;

export const HeaderText = css`
    max-width: 50%;
    padding-right: 16px;
`;

export const PopupButton = css`
    width: 100px;
    margin-left: 8px;
`;

export const ErrorsBlock = css`
    margin-top: 0px;
    margin-bottom: 0px;
    padding-left: 12px;
`;