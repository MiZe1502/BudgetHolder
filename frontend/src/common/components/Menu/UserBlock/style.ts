import {css} from 'emotion'

import {
    SideMinorPadding, FlexVert, FlexHorCenter
} from "../../../../common/styles/positioning";
import {Font716Black, Font312Black} from "../../../../common/styles/fonts";

import colors from "../../../../common/styles/colors";

const imageHeight = 44;

export const style = {
    SideMinorPadding,
    FlexVert,
    FlexHorCenter,

    Font716Black,
    Font312Black,

    Image: css`
        height: ${imageHeight}px;
        
        &:hover {
          cursor: pointer;
        }
    `,

    UserBlock: css`
        padding-top: 4px;
        background-color: ${colors.greenPrimary};
        border-radius: 50%;
        position: relative;
    `,

    Login: css`
        margin-right: 8px;
    `,

    PopupItem: css`
        cursor: pointer;
        padding: 4px 8px;
        height: 30px;
        background-color: ${colors.white};
        border-bottom: 1px solid ${colors.darkGray};
    
        &:hover {
          cursor: pointer;
          background-color: ${colors.greenLight};
        }
        
        &:first-of-type {
          border-top-left-radius: 4px;
          border-top-right-radius: 4px;
        }
        
        &:last-of-type {
          border: 0;
          border-bottom-left-radius: 4px;
          border-bottom-right-radius: 4px;
        }
    `
}