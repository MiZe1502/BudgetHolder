import { css } from 'emotion';
import colors from "../../styles/colors";

import { Font724Black, Font312Red } from "../../styles/fonts";
import { FlexHorCenter, Overflowed, FlexVert } from "../../styles/positioning";

export const style = {
    Font724Black,
    Font312Red,

    FlexHorCenter,
    Overflowed,
    FlexVert,

    Popup: css`
        position: fixed;
        margin: auto;
        width: 900px;
        height: auto;
        border-radius: 4px;
        background: ${colors.white};
        box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
    `,

    UtilBlock: css`
        width: 100%;
        height: 50px;
        background-color: ${colors.gray};
        padding: 16px;
        box-sizing: border-box;
    `,

    Header: css`
        cursor: grab;
        border-top-right-radius: 4px;
        border-top-left-radius: 4px;
        justify-content: space-between;
    `,

    Footer: css`
        border-bottom-right-radius: 4px;
        border-bottom-left-radius: 4px;
        height: auto;
        justify-content: flex-end;
    `,

    FooterWithErrors: css`
        justify-content: space-between;
    `,

    HeaderText: css`
        max-width: 50%;
        padding-right: 16px;
    `,

    PopupButton: css`
        width: 100px;
        margin-left: 8px;
    `,

    ErrorsBlock: css`
        margin-top: 0px;
        margin-bottom: 0px;
        padding-left: 12px;
    `
}