import { css } from 'emotion';
import colors from "../../../styles/colors";

import { FlexHorCenter, SideMainPadding } from "../../../styles/positioning";

export const style = {
    FlexHorCenter,
    SideMainPadding,

    MainMenu: css`
        width: 100%;
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.3);
        margin-bottom: 40px;
        background-color: ${colors.greenPrimary};
        position: fixed;
        z-index: 2;
    `,

    Wrapper: css`
        justify-content: space-between;
        width: 100%;
    `
}