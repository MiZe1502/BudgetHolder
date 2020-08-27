import { css } from 'emotion';
import colors from "../../../styles/colors";

import { Font716Black } from "../../../styles/fonts";
import { SideMainPadding, FlexHorCenter, Overflowed } from "../../../styles/positioning";

export const style = {
    Font716Black,

    SideMainPadding,
    FlexHorCenter,
    Overflowed,

    Row: css`
        min-height: 40px;
        padding-top: 8px;
        padding-bottom: 8px;
        align-items: flex-start;
    
        &:nth-of-type(2n) {
            background-color: ${colors.gray};
        }
    `,

    Cell: css`
        margin-right: 16px;
    `,

    PopupPadding: css`
        padding-left: 16px;
        padding-right: 16px;
    `
}