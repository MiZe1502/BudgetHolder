import {css} from 'emotion'

export {
    SideMinorPadding, FlexVert, FlexHorCenter
} from "../../../../common/styles/positioning";
export {Font716Black, Font312Black} from "../../../../common/styles/fonts";

import colors from "../../../../common/styles/colors";

const imageHeight = 44;

export const Image = css`
    height: ${imageHeight}px;
    
    &:hover {
      cursor: pointer;
    }
`;

export const UserBlock = css`
    padding-top: 4px;
    background-color: ${colors.greenPrimary};
    border-radius: 50%;
    position: relative;
`

export const Login = css`
    margin-right: 8px;
`;

export const PopupItem = css`
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
`;