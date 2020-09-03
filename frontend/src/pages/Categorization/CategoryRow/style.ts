import {css} from 'emotion'

import colors from "../../../common/styles/colors";

import {Font312Black} from "../../../common/styles/fonts";
import {
    SideMinorPadding, FlexHorCenter
} from "../../../common/styles/positioning";

export const style = {
    Font312Black,

    SideMinorPadding,
    FlexHorCenter,

    ArrowButton: css`
        margin-right: 8px;
    `,

    ArrowButtonExpanded: css`
        margin-right: 8px;
    
        & > svg {
            transform: rotate(90deg);
        }
    `,

    MarginedName: css`
        margin-left: 24px;
    `,

    CategoryLine: css`
        height: 40px;
        padding-right: 8px;
        border-bottom: 1px solid ${colors.darkGray};
        display: flex;
        justify-content: space-between;
    
        &:last-of-type {
            border-bottom: 0px;
        }
    `,

    CategoryLineExpanded: css`
      background-color: ${colors.gray};
    `
}