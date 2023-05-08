// Copyright 2023 Ryan McGuinness
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import { red } from '@mui/material/colors';


import { RDMModule } from './Model';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';


export default function ModuleCard(rdmModule: RDMModule) {
  return (
    <Card key={rdmModule.name} sx={{ maxWidth: 345 }}>
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
            {rdmModule.name.substring(0,1)}
          </Avatar>
        }
        action={
          <IconButton aria-label="settings">
            <MoreHorizIcon />
          </IconButton>
        }
        title={rdmModule.name}
        subheader={rdmModule.version}
      />
      <CardMedia
        component="img"
        height="194"
        image={rdmModule.banner}
        alt={rdmModule.bannerAlt}
      />
      <CardContent>
        <Typography variant="body2" color="text.secondary">
          {rdmModule.description}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="detail">
          <MoreHorizIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}
