import {css} from 'emotion'

import colors from "../../../common/styles/colors";
import {
    Font312DarkGray, Font312Black, Font312Red
} from "../../../common/styles/fonts";
import {
    SideMinorPadding, FlexVert, FlexHor, FlexHorCenter
} from "../../../common/styles/positioning";

const maxHeight = 400;

export const style = {
    Font312DarkGray,
    Font312Black,
    Font312Red,

    SideMinorPadding,
    FlexVert,
    FlexHor,
    FlexHorCenter,

    maxHeight,

    TextArea: css`
        min-height: 50px;
        height: 100%;
    `,

    ValidationBlock: css`
        width: 100%;;
    `,

    Form: css`
        padding-top: 8px;
        width: 100%;
    `,

    MainFieldsWrapper: css`
        border-bottom: 1px solid ${colors.darkGray};
        box-shadow: 0 2px 4px -2px rgba(0,0,0,0.25);
        padding-bottom: 16px;
        padding-left: 32px;
        padding-right: 32px;
    `,

    MinorFieldsWrapper: css`
        margin-right: 32px;
        margin-left: 32px;
        padding-top: 16px;
        padding-bottom: 16px;
        border-bottom: 1px solid ${colors.darkGray};
        
        &:last-of-type {
            border-bottom: none;
        }
    `,

    MainColumn: css`
        width: 50%;
    `,

    MinorColumn: css`
        width: 33.333%;
    `,

    NotLastColumn: css`
        margin-right: 32px;
    `,

    ButtonsWrapper: css`
        padding: 16px 32px;
        justify-content: space-between;
        box-shadow: 0 -2px 4px -2px rgba(0,0,0,0.25);
    `,

    ButtonsBlock: css`
        min-width: 30%;
        justify-content: flex-end;
    `,

    ButtonForm: css`
        margin-left: 8px;
    `,

    CounterBlock: css`
        border-radius: 4px;
        border: 1px solid ${colors.greenPrimary};
        padding: 4px;
        width: 24px;
        height: 24px;
        box-sizing: border-box;
        justify-content: center;
    `,

    GoodsItemControlBlock: css`
      justify-content: space-between;
    `
}
