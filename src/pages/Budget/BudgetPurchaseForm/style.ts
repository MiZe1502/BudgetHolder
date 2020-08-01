import {css} from 'emotion'

export {Font312DarkGray, Font312Black} from "../../../common/styles/fonts";
export {
    SideMinorPadding, FlexVert, FlexHor
} from "../../../common/styles/positioning";

import colors from "../../../common/styles/colors";

export const TextArea = css`
    min-height: 150px;
    height: 100%;
`;

export const Form = css`
    padding-top: 8px;
`;

export const MainFieldsWrapper = css`
    border-bottom: 1px solid ${colors.darkGray};
    padding-bottom: 32px;
`