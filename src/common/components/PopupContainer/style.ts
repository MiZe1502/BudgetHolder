import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font724Black } from "../../styles/fonts";
export { FlexHorCenter, Overflowed } from "../../styles/positioning";

export const Popup = css`
    position: fixed;
    margin: auto;
    width: 900px;
    height: 600px;
    border-radius: 4px;
    background: ${colors.white};
    box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
`;

export const Header = css`
    width: 100%;
    height: 50px;
    background-color: ${colors.gray};
    cursor: grab;
    border-top-right-radius: 4px;
    border-top-left-radius: 4px;
    padding: 16px;
    box-sizing: border-box;
    justify-content: space-between;
`;

export const HeaderText = css`
    max-width: 50%;
    padding-right: 16px;
`;