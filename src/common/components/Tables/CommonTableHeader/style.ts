import { css } from 'emotion';
import colors from "../../../styles/colors";

import { Font716Black } from "../../../styles/fonts";
import { SideMainPadding, FlexHorCenter } from "../../../styles/positioning";

export const style = {
    Font716Black,

    SideMainPadding,
    FlexHorCenter,

    Header: css`
        position: relative;
        min-height: 40px;
        background-color: ${colors.white};
        box-shadow: 0 2px 4px -2px rgba(0, 0, 0, 0.25);
    `,

    PopupPadding: css`
        padding-left: 16px;
        padding-right: 16px;
    `
}