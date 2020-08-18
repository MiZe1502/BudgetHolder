import { css } from 'emotion';
import colors from "../../../styles/colors";

export { Font716Black } from "../../../styles/fonts";
export { SideMainPadding, FlexHorCenter } from "../../../styles/positioning";

export const Header = css`
    position: relative;
    min-height: 40px;
    background-color: ${colors.white};
    box-shadow: 0 2px 4px -2px rgba(0, 0, 0, 0.25);
`;

export let PopupPadding = css`
    padding-left: 16px;
    padding-right: 16px;
`;