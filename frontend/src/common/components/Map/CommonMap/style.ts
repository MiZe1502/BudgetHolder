import { css } from 'emotion';
import colors from "../../../styles/colors";

import { SideMainPadding, FlexVert, SectionBottomMargin } from "../../../styles/positioning";

export const style = {
    SideMainPadding,
    FlexVert,
    SectionBottomMargin,

    MapContainerWrapper: css`
        height: 700px;
        margin-top: 8px;
        border-radius: 4px;
        border: 1px solid ${colors.darkGray};
        padding: 1px;
    `
}