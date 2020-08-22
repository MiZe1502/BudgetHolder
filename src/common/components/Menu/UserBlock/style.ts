import {css} from 'emotion'

export {
    SideMinorPadding, FlexVert
} from "../../../../common/styles/positioning";

import colors from "../../../../common/styles/colors";

export const Image = css`
    width: 44px;
    height: 44px;
    
    &:hover {
      cursor: pointer;
    }
`;

export const UserBlock = css`
    background-color: ${colors.greenPrimary};
`