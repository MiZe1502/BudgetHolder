import { css } from 'emotion'

import { Font312DarkGray, Font312Black } from "../../../common/styles/fonts";
import { SideMinorPadding, FlexVert } from "../../../common/styles/positioning";

export const style = {
    Font312DarkGray,
    Font312Black,

    SideMinorPadding,
    FlexVert,

    TextArea: css`
      min-height: 150px;
    `,

    Form: css`
        padding-top: 8px;
        padding-bottom: 8px;
    `
}